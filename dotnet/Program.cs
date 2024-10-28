using BookStoreApi.Models;
using BookStoreApi.Services;
using Examples.AspNetCore;
using OpenTelemetry.Exporter;
using OpenTelemetry.Instrumentation.AspNetCore;
using OpenTelemetry.Logs;
using OpenTelemetry.Metrics;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using System.Reflection.PortableExecutable;

var appBuilder = WebApplication.CreateBuilder(args);

const string serviceName = "bookstore-api";

var otlpHostName = Environment.GetEnvironmentVariable("OTLP_HOSTNAME") ?? "localhost";

appBuilder.Services.AddSingleton<Instrumentation>();

appBuilder.Logging.ClearProviders();

appBuilder.Services.AddOpenTelemetry()
    .ConfigureResource(r => r
        .AddService(
            serviceName: serviceName,
            serviceVersion: typeof(Program).Assembly.GetName().Version?.ToString() ?? "unknown",
            serviceInstanceId: Environment.MachineName))
    .WithTracing(builder =>
    {
        builder
            .AddSource(Instrumentation.ActivitySourceName)
            .SetSampler(new AlwaysOnSampler())
            .AddHttpClientInstrumentation()
            .AddAspNetCoreInstrumentation();

        appBuilder.Services.Configure<AspNetCoreTraceInstrumentationOptions>(appBuilder.Configuration.GetSection("AspNetCoreInstrumentation"));

        builder.AddOtlpExporter(otlpOptions =>
        {
            otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
        }).AddConsoleExporter();
    })
    .WithMetrics(builder =>
    {
        builder
            .AddMeter(Instrumentation.MeterName)
            .SetExemplarFilter(ExemplarFilterType.TraceBased)
            .AddRuntimeInstrumentation()
            .AddHttpClientInstrumentation()
            .AddAspNetCoreInstrumentation();

        builder.AddOtlpExporter(otlpOptions =>
        {
            otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
        }).AddConsoleExporter()
        .AddPrometheusExporter();
    }).WithLogging(builder =>
    {
        builder.AddOtlpExporter(otlpOptions =>
        {
            otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
        }).AddConsoleExporter();
    });

appBuilder.Services.Configure<BookStoreDatabaseSettings>(
  appBuilder.Configuration.GetSection("BookStoreDatabase"));

appBuilder.Services.AddSingleton<BooksService>();

appBuilder.Services.AddControllers()
  .AddJsonOptions(
    options => options.JsonSerializerOptions.PropertyNamingPolicy = null);

appBuilder.Services.AddEndpointsApiExplorer();
appBuilder.Services.AddSwaggerGen();

var app = appBuilder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.UseOpenTelemetryPrometheusScrapingEndpoint();

app.UseAuthorization();
app.MapControllers();
app.Run();
