using DotnetAppBlueprint;
using OpenTelemetry.Instrumentation.AspNetCore;
using OpenTelemetry.Logs;
using OpenTelemetry.Metrics;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace Microsoft.Extensions.DependencyInjection;

internal static class OpenTelemetryExtension
{
    const string serviceName = "AppDotNet";

    public static IServiceCollection AddTelemetry(this IServiceCollection services, IConfiguration configuration)
    {

        var otlpHostName = configuration.GetValue<string>("OTLP_HOSTNAME");

        services.AddOpenTelemetry()
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

                    services.Configure<AspNetCoreTraceInstrumentationOptions>(configuration.GetSection("AspNetCoreInstrumentation"));

                    builder.AddOtlpExporter(otlpOptions =>
                    {
                        otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
                    }).AddConsoleExporter();
                })

                .WithMetrics(builder =>
                {
                    builder
                        .SetExemplarFilter(ExemplarFilterType.TraceBased)
                        .AddRuntimeInstrumentation()
                        .AddHttpClientInstrumentation()
                        .AddAspNetCoreInstrumentation();

                    builder.AddOtlpExporter(otlpOptions =>
                    {
                        otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
                    }).AddConsoleExporter()
                    .AddPrometheusExporter();
                })

                .WithLogging(builder =>
                {
                    builder.AddOtlpExporter(otlpOptions =>
                    {
                        otlpOptions.Endpoint = new Uri($"http://{otlpHostName}:4317");
                    }).AddConsoleExporter();
                });

        return services;
    }
}