import { BookRecord } from '../types/Book'
import { tsundokuPrice } from './tsundoku'
import dayjs from 'dayjs'

export const comparePrice = (
  isDesc: boolean
): ((a: BookRecord, b: BookRecord) => number) => (
  a: BookRecord,
  b: BookRecord
): number => {
  return (a.price - b.price) * (isDesc ? -1 : 1)
}

export const compareTsundokuPrice = (
  isDesc: boolean
): ((a: BookRecord, b: BookRecord) => number) => (
  a: BookRecord,
  b: BookRecord
): number => {
  return (
    (tsundokuPrice(a.readPages, a.totalPages, a.price) -
      tsundokuPrice(b.readPages, b.totalPages, b.price)) *
    (isDesc ? -1 : 1)
  )
}

export const compareUpdatedAt = (
  isDesc: boolean
): ((a: BookRecord, b: BookRecord) => number) => (
  a: BookRecord,
  b: BookRecord
): number => {
  return (
    (dayjs(b.updatedAt).unix() - dayjs(a.updatedAt).unix()) * (isDesc ? -1 : 1)
  )
}

export const compareTotalPages = (
  isDesc: boolean
): ((a: BookRecord, b: BookRecord) => number) => (
  a: BookRecord,
  b: BookRecord
): number => {
  return (a.totalPages - b.totalPages) * (isDesc ? -1 : 1)
}
