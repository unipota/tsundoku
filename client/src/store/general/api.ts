import axios, { AxiosResponse } from 'axios'
import { BookRecord } from '@/types/Book'

const client = axios.create({
  baseURL: 'https://test.tsun-doku.app/',
  headers: { 'Access-Control-Allow-Origin': '*' }
})

const api = {
  getMyBooks(): Promise<AxiosResponse> {
    return client.get('api/books')
  },
  addNewBook(book: BookRecord): Promise<AxiosResponse> {
    return client.post('api/books', book)
  },
  getBookDetail(id: number): Promise<AxiosResponse> {
    return client.get(`api/books/${id}`)
  },
  updateBook(id: number, book: BookRecord): Promise<AxiosResponse> {
    return client.put(`api/books/${id}`, book)
  },
  searchBooksByISBN(isbn: string): Promise<AxiosResponse> {
    return client.get('api/search/isbn', { params: { isbn } })
  },
  searchBooks(search: string): Promise<AxiosResponse> {
    return client.get('api/search', { params: { search } })
  }
}

export default api
