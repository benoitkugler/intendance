import levenshtein from "js-levenshtein";
import { Horaire, HoraireLabels } from "@/logic/types";
const MAX_DIST_LEVENSHTEIN = 5;

const N = Object.keys(Horaire).length;
const computeColor = (i: number) =>
  `rgb(${200 * (1 - i / N)},${(100 * i) / N},50)`;
export const HorairesColors: { [key in Horaire]: string } = {
  [Horaire.PetitDejeuner]: computeColor(Horaire.PetitDejeuner),
  [Horaire.Midi]: computeColor(Horaire.Midi),
  [Horaire.Gouter]: computeColor(Horaire.Gouter),
  [Horaire.Diner]: computeColor(Horaire.Diner),
  [Horaire.Cinquieme]: computeColor(Horaire.Cinquieme)
};

export const HorairesIcons = {
  [Horaire.PetitDejeuner]: "food-croissant",
  [Horaire.Midi]: "pasta",
  [Horaire.Gouter]: "cupcake",
  [Horaire.Diner]: "bowl-mix",
  [Horaire.Cinquieme]: "glass-mug-variant"
};

export const Days = [
  "Lundi",
  "Mardi",
  "Mercredi",
  "Jeudi",
  "Vendredi",
  "Samedi",
  "Dimanche"
];

export function sortByText<T extends { text: string }>(l: T[]) {
  return l.sort((a, b) => Number(a.text < b.text));
}

export type Crible = { [key: number]: boolean };

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

// Implémente le délai d'un requête
export class Debounce {
  private timerId: null | number = null;
  private job: () => void;
  private delay = 500; // default ms

  constructor(job: () => void, delay?: number) {
    this.job = job;
    if (delay !== undefined) {
      this.delay = delay;
    }
  }

  // lance le job après un delai
  delayJob() {
    // un job a déjà était proposé :
    // on l'annule et on reset le timer
    if (this.timerId != null) {
      clearTimeout(this.timerId);
      this.timerId = null;
    }

    // on sauvegarde le lancement du job
    this.timerId = setTimeout(this.job, this.delay);
  }
}

/** permute les éléments de la liste en mettant le premier à la fin */
export function cycleDays<T>(l: T[]): T[] {
  const out: T[] = [];
  l.map((_, i) => {
    const index = (i - 1 + 7) % 7; // recule de 1
    out.push(l[index]);
  });
  return out;
}
