using BookStoreApi.Models;
using BookStoreApi.Services;
using Examples.AspNetCore;
using OpenTelemetry.Logs;
using OpenTelemetry.Metrics;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using System.Reflection.PortableExecutable;

var builder = WebApplication.CreateBuilder(args);

const string serviceName = "bookstore-api";

var otlpHostName = Environment.GetEnvironmentVariable("OTLP_HOSTNAME") ?? "localhost";

builder.Logging.ClearProviders()
    .AddOpenTelemetry(options =>
    {
        options.IncludeFormattedMessage = true;
        options.IncludeScopes = true;
        options.SetResourceBuilder(ResourceBuilder.CreateDefault().AddService(serviceName))
        .AddOtlpExporter(options => options.Endpoint = new Uri($"http://{otlpHostName}:4317"))
        .AddConsoleExporter();
    });

builder.Services.AddOpenTelemetry()
      .ConfigureResource(resource => resource.AddService(serviceName))
      .WithTracing(tracing => tracing
          .AddAspNetCoreInstrumentation()
          .AddOtlpExporter(options => options.Endpoint = new Uri($"http://{otlpHostName}:4317"))
          .AddConsoleExporter())
      .WithMetrics(metrics => metrics
          .AddAspNetCoreInstrumentation()
          .AddOtlpExporter(options => options.Endpoint = new Uri($"http://{otlpHostName}:4317"))
          .AddConsoleExporter());

builder.Services.AddSingleton<Instrumentation>();

builder.Services.Configure<BookStoreDatabaseSettings>(
  builder.Configuration.GetSection("BookStoreDatabase"));

builder.Services.AddSingleton<BooksService>();

builder.Services.AddControllers()
  .AddJsonOptions(
    options => options.JsonSerializerOptions.PropertyNamingPolicy = null);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.UseAuthorization();
app.MapControllers();
app.Run();
