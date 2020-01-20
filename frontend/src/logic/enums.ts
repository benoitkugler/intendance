interface enumItem {
  value: string;
  text: string;
}

function fmt(enums: enumItem[], value: string) {
  const data = enums.find(v => v.value == value);
  if (data) return data.text;
  return "";
}

export const HoraireFields = {
  PetitDejeuner: "matin",
  Midi: "midi",
  Gouter: "gouter",
  Diner: "diner",
  Cinquieme: "cinquieme"
};
export const Horaires = [
  { value: HoraireFields.PetitDejeuner, text: "Petit déjeuner" },
  { value: HoraireFields.Midi, text: "Midi" },
  { value: HoraireFields.Gouter, text: "Goûter" },
  { value: HoraireFields.Diner, text: "Dinner" },
  { value: HoraireFields.Cinquieme, text: "Cinquième" }
];
export const fmtHoraire = (v: string) => fmt(Horaires, v);

export const UniteFields = { Litres: "L", Kilos: "Kg", Piece: "P" };
export const Unites = [
  { value: UniteFields.Litres, text: "Litres" },
  { value: UniteFields.Kilos, text: "Kilos" },
  { value: UniteFields.Piece, text: "Pièce(s)" }
];
export const fmtUnite = (v: string) => fmt(Unites, v);
