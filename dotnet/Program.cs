using System.Text.Json;
using Configuration;
using DotnetAppBlueprint;
using Microsoft.AspNetCore.Http.Json;

var builder = WebApplication.CreateBuilder(args);
builder.WebHost.CaptureStartupErrors(true);

builder.Services.AddSingleton<Instrumentation>();

builder.Logging.ClearProviders();

builder.Services.AddTelemetry(builder.Configuration);

builder.Services.Configure<JsonOptions>(options =>
{
    options.SerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.CamelCase;
});

builder.Services.AddOptions<DatabaseConfiguration>().Bind(builder.Configuration.GetSection("MongoDatabase"));

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.UseOpenTelemetryPrometheusScrapingEndpoint();

app.MapRoutes();

await app.RunAsync();

public partial class Program { }