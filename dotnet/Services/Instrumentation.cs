namespace Examples.AspNetCore;

using System.Diagnostics;
using System.Diagnostics.Metrics;

public class Instrumentation : IDisposable
{
    internal const string ActivitySourceName = "Bookstore";
    internal const string MeterName = "Bookstore_meter";
    private readonly Meter meter;

    public Instrumentation()
    {
        string? version = typeof(Instrumentation).Assembly.GetName().Version?.ToString();
        this.ActivitySource = new ActivitySource(ActivitySourceName, version);
        this.meter = new Meter(MeterName, version);
        this.BooksCounter = this.meter.CreateCounter<long>("books.count", description: "The number of books");
    }

    public ActivitySource ActivitySource { get; }

    public Counter<long> BooksCounter { get; }

    public void Dispose()
    {
        this.ActivitySource.Dispose();
        this.meter.Dispose();
    }
}