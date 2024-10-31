namespace DotnetAppBlueprint;

internal static class Api
{
    internal static void MapRoutes(this WebApplication app)
    {
        var group = app.MapGroup("/api").WithTags("api");

        group.MapGet("/long_run", ApiFunctions.LongRun);

        group.MapGet("/short_run", ApiFunctions.ShortRun);

        group.MapGet("/database_run", ApiFunctions.DatabaseRun);

        group.MapGet("failed_run", ApiFunctions.FailedRun);

        app.MapGet("/", () => new { data = "hello world" }).WithTags("api");
    }
}
