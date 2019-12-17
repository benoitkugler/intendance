export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;
