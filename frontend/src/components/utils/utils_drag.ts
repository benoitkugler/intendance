/** Helpers for type-safe drag and drop operations */

import { RepasComplet, MenuComplet } from "@/logic/api";

interface GroupeDragData {
  idGroupe: number;
  repas: RepasComplet;
}

/** Définit les types de drag and drop possibles. */
export enum DragKind {
  Repas = "repas",
  Menu = "menu",
  Groupe = "groupe",
  IdIngredient = "id-ingredient",
  IdRecette = "id-recette",
}

interface dragTypes {
  [DragKind.Repas]: RepasComplet;
  [DragKind.Menu]: MenuComplet;
  [DragKind.IdIngredient]: number;
  [DragKind.IdRecette]: number;
  [DragKind.Groupe]: GroupeDragData;
}

export function setDragData<T extends keyof dragTypes>(
  transfer: DataTransfer,
  kind: T,
  data: dragTypes[T]
) {
  transfer.setData(kind, JSON.stringify(data));
}

export function getDragData<T extends keyof dragTypes>(
  transfer: DataTransfer,
  kind: T
): dragTypes[T] {
  return JSON.parse(transfer.getData(kind));
}
