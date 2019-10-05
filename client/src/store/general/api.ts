import axios, { AxiosResponse } from 'axios'
import { BookSimpleRecord, BookRecord, BookStats } from '@/types/Book'

const isDev = process.env['NODE_ENV'] === 'development'

const client = axios.create({
  baseURL: isDev ? '/' : 'https://test.tsun-doku.app/',
  headers: { 'Access-Control-Allow-Origin': '*' }
})

const api = {
  getMyBooks(): Promise<AxiosResponse<BookRecord[]>> {
    return client.get('api/books')
  },
  addNewBook(book: BookSimpleRecord): Promise<AxiosResponse> {
    return client.post('api/books', book)
  },
  getBookDetail(id: string): Promise<AxiosResponse> {
    return client.get(`api/books/${id}`)
  },
  updateBook(id: string, book: BookRecord): Promise<AxiosResponse<BookRecord>> {
    return client.put(`api/books/${id}`, book)
  },
  deleteBook(id: string): Promise<AxiosResponse> {
    return client.delete(`api/books/${id}`)
  },
  searchBooksByISBN(isbn: string): Promise<AxiosResponse<BookRecord[]>> {
    return client.get('api/search/isbn', { params: { isbn } })
  },
  searchBooks(search: string): Promise<AxiosResponse<BookRecord[]>> {
    return client.get('api/search', { params: { search } })
  },
  getAllBookStats(): Promise<AxiosResponse<BookStats[]>> {
    return client.get('api/books/statistics')
  }
}

export default api
