using System.Diagnostics;
using MongoDB.Bson;
using MongoDB.Driver;
using OpenTelemetry.Trace;

namespace DotnetAppBlueprint;

public record DataDto(int Data);

internal static class Query
{
    public static async Task<List<DataDto>> ReadData(
        this IMongoCollection<BsonDocument> collection,
        FilterDefinition<BsonDocument> filter,
        Instrumentation instrumentation,
        ILogger logger)
    {
        using var activity = instrumentation.ActivitySource.StartActivity("ReadData")!;
        activity.AddEvent(new("ReadData started"));

        try
        {
            var result = await collection.Find(filter).ToListAsync();

            activity.SetStatus(ActivityStatusCode.Ok, "ok");
            activity.AddEvent(new("ReadData done"));
            return result.Select(r => new DataDto(r.GetValue("data").AsInt32)).ToList();
        }
        catch (Exception ex)
        {
            logger.LogError(ex, "{Message}", ex.Message);

            activity.SetStatus(ActivityStatusCode.Error, "UnexpectedError");
            activity.RecordException(ex);
            throw;
        }
    }
}
