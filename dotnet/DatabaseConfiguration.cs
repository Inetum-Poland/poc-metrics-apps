namespace Configuration;

public record DatabaseConfiguration
{
    public required string ConnectionString { get; init; }
    public required string DatabaseName { get; init; }
}