import { Horaire } from "@/logic/types";

export const UnitePiece = "P";

export const Unites = [
  { text: "Litres", value: "L" },
  { text: "Kilos", value: "Kg" },
  { text: "Pièces", value: UnitePiece }
];

export const Horaires: {
  text: string;
  value: string;
}[] = [
    { text: "Petit déjeuner", value: "petit_dej" },
    { text: "Midi", value: "midi" },
    { text: "Goûter", value: "gouter" },
    { text: "Diner", value: "diner" },
    { text: "Cinquième", value: "cinquieme" }
  ];
