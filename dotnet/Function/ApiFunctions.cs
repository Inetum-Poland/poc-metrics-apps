using System.Diagnostics;
using Configuration;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;
using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;
using OpenTelemetry.Trace;

namespace DotnetAppBlueprint;

internal static class ApiFunctions
{
    public static async Task<IResult> LongRun([FromServices]Instrumentation instrumentation)
    {

        using var activity = instrumentation.ActivitySource.StartActivity("LongRun")!;
        activity.AddEvent(new("LongRun started"));

        await Functions.Add(instrumentation.ActivitySource, 1, 2);
        await Functions.Subtract(instrumentation.ActivitySource, 1, 2);
        await Functions.Multiply(instrumentation.ActivitySource, 1, 2);
        await Functions.Divide(instrumentation.ActivitySource, 1, 2);

        activity.SetStatus(ActivityStatusCode.Ok, "ok");
        activity.AddEvent(new("LongRun done"));
        return Results.Ok(new { data = "ok" });
    }

    public static async Task<IResult> ShortRun([FromServices]Instrumentation instrumentation)
    {
        using var activity = instrumentation.ActivitySource.StartActivity("ShortRun")!;
        activity.AddEvent(new("ShortRun started"));

        await Task.Delay(100);

        activity.SetStatus(ActivityStatusCode.Ok, "ok");
        activity.AddEvent(new("ShortRun done"));
        return Results.Ok(new { data = "ok" });
    }

    public static async Task<List<DataDto>> DatabaseRun([FromServices]Instrumentation instrumentation, [FromServices]ILogger<Program> logger, [FromServices]IOptions<DatabaseConfiguration> configuration)
    {
        using var activity = instrumentation.ActivitySource.StartActivity("DatabaseRun")!;
        activity.AddEvent(new("DatabaseRun started"));

        var filter = Builders<BsonDocument>.Filter.Gte("data", 0);

        var collection = await Database.Connect(instrumentation, logger, configuration.Value);

        activity.AddEvent(new("collection.Find"));

        BsonDocument bsonDocument = filter.Render(new(
            BsonDocumentSerializer.Instance, // The serializer for BsonDocument
            BsonSerializer.SerializerRegistry // Default serializer registry
        ));

        activity.SetTag("query", bsonDocument.ToJson());

        try
        {
            var result = await collection.ReadData(filter, instrumentation, logger);

            activity.SetStatus(ActivityStatusCode.Ok, "ok");
            activity.AddEvent(new("DatabaseRun done"));
            return result;
        }
        catch (Exception ex)
        {
            logger.LogError(ex, "{Message}", ex.Message);

            activity.AddEvent(new("DatabaseRun error"));
            activity.SetStatus(ActivityStatusCode.Error, ex.Message);
            activity.RecordException(ex);
            throw;
        }
    }

    public static async Task<IResult> FailedRun([FromServices]Instrumentation instrumentation, [FromServices]ILogger<Program> logger)
    {
        using var activity = instrumentation.ActivitySource.StartActivity("FailedRun")!;
        activity.AddEvent(new("FailedRun started"));

        await Task.Delay(100);

        logger.LogError("Failed");

        activity.SetStatus(ActivityStatusCode.Error, "UnexpectedError");
        activity.RecordException(new("databaseRun done"));
        return Results.Json(new { data = "nok" }, statusCode: 500);
    }
}