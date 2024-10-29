using System.Diagnostics;

namespace DotnetAppBlueprint;

internal static class Functions
{
    private static int RandomRange(int min, int max) => Random.Shared.Next(min, max) + min;

    private static async Task<int> PerformOperation(ActivitySource ctx, string operation, int a, int b, Func<int, int, int> opFunc)
    {
        using var activity = ctx.StartActivity(operation)!;
        activity.AddEvent(new($"{operation} started"));

        await Task.Delay(RandomRange(20, 800));
        var result = opFunc(a, b);

        activity.SetStatus(ActivityStatusCode.Ok, "ok");
        activity.AddEvent(new($"{operation} done"));

        return result;
    }

    public static Task<int> Add(ActivitySource ctx, int a, int b) => PerformOperation(ctx, "Add", a, b, (a, b) => a + b);

    public static Task<int> Subtract(ActivitySource ctx, int a, int b) => PerformOperation(ctx, "Subtract", a, b, (a, b) => a - b);

    public static Task<int> Multiply(ActivitySource ctx, int a, int b) => PerformOperation(ctx, "Multiply", a, b, (a, b) => a * b);

    public static Task<int> Divide(ActivitySource ctx, int a, int b) => PerformOperation(ctx, "Divide", a, b, (a, b) => a / b);
}