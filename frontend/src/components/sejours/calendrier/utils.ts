export interface DateTime {
  date: string;
  time: string;
}

export function toDateVuetify(d: Date) {
  return d.toISOString().substr(0, 10);
}
