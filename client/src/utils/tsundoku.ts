export const tsundokuPrice = (
  readPages: number,
  totalPages: number,
  price: number
): number => {
  if (totalPages === 0) {
    return price
  }
  return Math.round((1 - readPages / totalPages) * price)
}
