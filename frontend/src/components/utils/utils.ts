import { Horaires } from "@/logic/enums";
import levenshtein from "js-levenshtein";
const MAX_DIST_LEVENSHTEIN = 5;

const N = Horaires.length;

export const HorairesColors: { [key: string]: string } = {};
Horaires.forEach((h, i) => {
  HorairesColors[h.value] = `rgb(100,${200 * (1 - i / N)}, ${(255 * i) / N}`;
});

export function sortByText<T extends { text: string }>(l: T[]) {
  return l.sort((a, b) => Number(a.text < b.text));
}

// Renvoie `true` si les deux tableaux sont égaux, vus comme ensembles.
export function compareArrays<T>(a: T[], b: T[]) {
  const sa = new Set(a);
  const sb = new Set(b);
  if (sa.size != sb.size) return false;
  for (const elem of sa) {
    if (!sb.has(elem)) return false;
  }
  return true;
}

// Renvoie un prédicat correspondant à la recherche de `search`
export function searchFunction(search: string) {
  // recherche vide : tout match
  if (!search) return (item: string) => true;

  let filterNom: (nom: string) => boolean;

  // on essaie en mode RegExp
  try {
    const s = new RegExp(search, "i");
    filterNom = nom => s.test(nom);
  } catch {
    const sl = search.toLowerCase();
    filterNom = (nom: string) => nom.includes(sl);
  }

  return function(item: string) {
    item = item.toLowerCase();

    // on essaie une recherche 'exacte'
    if (filterNom(item)) return true;

    // puis on relâche à un critère approché
    return levenshtein(item, search) <= MAX_DIST_LEVENSHTEIN;
  };
}
