import { Horaires } from "@/logic/enums";

const N = Horaires.length;

export const HorairesColors: { [key: string]: string } = {};
Horaires.forEach((h, i) => {
  HorairesColors[h.value] = `rgb(100,${200 * (1 - i / N)}, ${(255 * i) / N}`;
});
