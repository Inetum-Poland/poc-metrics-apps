using System.Diagnostics;

namespace DotnetAppBlueprint;

internal class Instrumentation : IDisposable
{
    internal const string ActivitySourceName = "dotNet-app";

    public Instrumentation()
    {
        string? version = typeof(Instrumentation).Assembly.GetName().Version?.ToString();
        ActivitySource = new ActivitySource(ActivitySourceName, version);
    }

    public ActivitySource ActivitySource { get; }

    public void Dispose()
    {
        Dispose(true);
        GC.SuppressFinalize(this);
    }

    protected virtual void Dispose(bool disposing)
    {
        if (disposing)
        {
            ActivitySource.Dispose();
        }
    }

    ~Instrumentation()
    {
        Dispose(false);
    }
}