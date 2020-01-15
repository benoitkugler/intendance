import { Horaire } from "@/logic/types";

export const UnitePiece = "P";

export const Unites = [
  { text: "Litres", value: "L" },
  { text: "Kilos", value: "Kg" },
  { text: "Pièces", value: UnitePiece }
];

export const Horaires: {
  text: string;
  value: Horaire;
}[] = [
  { text: "Petit déjeuner", value: { heure: 8, minute: 0 } },
  { text: "Midi", value: { heure: 12, minute: 0 } },
  { text: "Goûter", value: { heure: 16, minute: 30 } },
  { text: "Diner", value: { heure: 19, minute: 0 } },
  { text: "Cinquième", value: { heure: 21, minute: 30 } }
];
