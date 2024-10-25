using BookStoreApi.Models;
using BookStoreApi.Services;
using Examples.AspNetCore;
using Microsoft.AspNetCore.Mvc;
using System.Diagnostics.Metrics;
using System.Diagnostics;
using static System.Reflection.Metadata.BlobBuilder;

namespace BookStoreApi.Controllers;

[ApiController]
[Route("api/[controller]")]
public class BooksController : ControllerBase
{
    private readonly BooksService _booksService;
    private readonly ILogger<BooksController> _logger;
    private readonly ActivitySource _activitySource;
    private readonly Counter<long> _booksCounter;

    public BooksController(BooksService booksService, ILogger<BooksController> logger, Instrumentation instrumentation)
    {
        _booksService = booksService;

        _logger = logger;
        _activitySource = instrumentation.ActivitySource;
        _booksCounter = instrumentation.BooksCounter;
    }

    [HttpGet]
    public async Task<List<Book>> Get()
    {
        using var activity = _activitySource.StartActivity("get books");

        var books = await _booksService.GetAsync();

        _booksCounter.Add(books.Count());

        _logger.LogInformation(
            "Books {count}: {books}",
            books.Count,
            books);

        return books;
    }

    [HttpGet("{id:length(24)}")]
    public async Task<ActionResult<Book>> Get(string id)
    {
        using var activity = _activitySource.StartActivity($"get book {id}");
        var book = await _booksService.GetAsync(id);

        if (book is null)
        {
            return NotFound();
        }

        _logger.LogInformation(
            "Book: {book}",
            book);

        return book;
    }

    [HttpPost]
    public async Task<IActionResult> Post(Book newBook)
    {
        using var activity = _activitySource.StartActivity($"create book {newBook.Author}: {newBook.BookName}");
        await _booksService.CreateAsync(newBook);

        _logger.LogInformation(
            "Book created: {book}",
            newBook);

        return CreatedAtAction(nameof(Get), new { id = newBook.Id }, newBook);
    }

    [HttpPut("{id:length(24)}")]
    public async Task<IActionResult> Update(string id, Book updatedBook)
    {
        using var activity = _activitySource.StartActivity($"update book {id}");
        var book = await _booksService.GetAsync(id);

        if (book is null)
        {
            return NotFound();
        }

        updatedBook.Id = book.Id;

        await _booksService.UpdateAsync(id, updatedBook);

        _logger.LogInformation(
            "Book updated: {book}",
            updatedBook);

        return NoContent();
    }

    [HttpDelete("{id:length(24)}")]
    public async Task<IActionResult> Delete(string id)
    {
        using var activity = _activitySource.StartActivity($"delete book {id}");
        var book = await _booksService.GetAsync(id);

        if (book is null)
        {
            return NotFound();
        }

        await _booksService.RemoveAsync(id);

        _logger.LogInformation(
            "Book deleted: {book}",
            book);

        return NoContent();
    }
}