using System.Diagnostics;
using Configuration;
using MongoDB.Bson;
using MongoDB.Driver;
using OpenTelemetry.Trace;

namespace DotnetAppBlueprint;

internal static class Database
{
    public static async Task<IMongoCollection<BsonDocument>> Connect(Instrumentation instrumentation, ILogger logger, DatabaseConfiguration configuration)
    {
        using var activity = instrumentation.ActivitySource.StartActivity("Connect")!;
        activity.AddEvent(new("Connect started"));

        try
        {
            var client = new MongoClient(configuration.ConnectionString);

            var database = client.GetDatabase(configuration.DatabaseName);
            
            await database.RunCommandAsync<BsonDocument>(new BsonDocument { { "ping", 1 } });

            var collection = database.GetCollection<BsonDocument>("Data");

            activity.SetStatus(ActivityStatusCode.Ok, "ok");
            activity.AddEvent(new("Connect done"));

            return collection;
        }
        catch (Exception ex)
        {
            logger.LogError(ex, "{Message}", ex.Message);

            activity.AddEvent(new("Connect error"));
            activity.SetStatus(ActivityStatusCode.Error, ex.Message);
            activity.RecordException(ex);
            throw;
        }
    }
}
