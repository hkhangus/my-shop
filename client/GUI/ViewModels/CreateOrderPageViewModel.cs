﻿// CreateOrderViewModel.cs
using CommunityToolkit.Mvvm.ComponentModel;
using Entity;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using ThreeLayerContract;

namespace GUI.ViewModels
{
    public class CreateOrderViewModel : ObservableObject
    {
        private ObservableCollection<BookWithSelection> _booksWithSelection;
        private Dictionary<string, IBus> _bus = BusInstance._bus;

        public CreateOrderViewModel()
        {
            var configuration = new Dictionary<string, string> { { "size", int.MaxValue.ToString() } };
            List<Book> availableBooks = new List<Book>(_bus["Book"].Get(configuration));

            _booksWithSelection = new ObservableCollection<BookWithSelection>();
            foreach (var book in availableBooks)
            {
                _booksWithSelection.Add(ConvertToBookWithSelection(book));
            }

        }

        public ObservableCollection<BookWithSelection> BookWithSelections
        {
            get => _booksWithSelection;
            set => SetProperty(ref _booksWithSelection, value);
        }

        private BookWithSelection ConvertToBookWithSelection(Book book)
        {
            return new BookWithSelection(book);
        }

        public ObservableCollection<BookWithSelection> BooksWithSelection
        {
            get => _booksWithSelection;
            set => SetProperty(ref _booksWithSelection, value);
        }

        public decimal TotalAmount => BooksWithSelection.Where(b => b.IsSelected).Sum(b => b.Subtotal);
    }

    public class BookWithSelection : ObservableObject
    {
        private bool _isSelected;
        private Book _book;
        private int _quantity;

        public BookWithSelection(Book book)
        {
            _book = book;
        }

        public bool IsSelected
        {
            get => _isSelected;
            set
            {
                if (SetProperty(ref _isSelected, value))
                {
                    OnPropertyChanged(nameof(Subtotal));
                    OnPropertyChanged(nameof(TotalAmount));
                }
            }
        }

        public string Name => _book.name;
        public string Author => _book.author;
        public decimal Price => _book.price;
        public int TotalSold => _book.total_sold;

        public int Quantity
        {
            get => _quantity;
            set
            {
                if (SetProperty(ref _quantity, value))
                {
                    OnPropertyChanged(nameof(Subtotal));
                    OnPropertyChanged(nameof(TotalAmount));
                }
            }
        }

        public decimal Subtotal => IsSelected ? _book.price * Quantity : 0;

        public decimal TotalAmount => IsSelected ? _book.price * Quantity : 0;
    }

}
