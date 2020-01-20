import { RepasWithGroupe } from "@/logic/types";

export function toDateVuetify(d: Date) {
  return d.toISOString().substr(0, 10);
}

export function formatNbOffset(repas: RepasWithGroupe) {
  const n = repas.offset_personnes;
  if (n != 0) {
    return `${n > 0 ? "+" : ""}${n} pers.`;
  }
  return "";
}
