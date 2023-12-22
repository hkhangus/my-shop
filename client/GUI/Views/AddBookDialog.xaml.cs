using Entity;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;

// To learn more about WinUI, the WinUI project structure,
// and more about our project templates, see: http://aka.ms/winui-project-info.

namespace GUI.Views
{
    /// <summary>
    /// An empty window that can be used on its own or navigated to within a Frame.
    /// </summary>
    public sealed partial class AddBookDialog : Window
    {
        List<BookCategory> _categories;
        public Book _newBook;
        public AddBookDialog()
        {
            this.InitializeComponent();
            LoadCategories();

            _newBook = new Book();

            bookForm.DataContext = _newBook;
        }

        private void LoadCategories()
        {
            _categories = new List<BookCategory>() {
                new BookCategory() { Id=1, Name="Novel"},
                new BookCategory() { Id=2, Name="Manga"}
            };

            categoriesComboBox.ItemsSource = _categories;
        }

        private void checkPrice(object sender, TextChangedEventArgs e)
        {
            
        }
    }
}
