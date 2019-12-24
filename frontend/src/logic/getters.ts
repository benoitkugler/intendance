import { Menu, Recette } from "./types";
import { D } from "./controller";
import { IngredientOptions } from "./types2";

export class G {
  static getAllIngredients(): IngredientOptions[] {
    return Object.values(D.ingredients).map(ing => {
      return { ingredient: ing };
    });
  }

  static getMenuRecettes(menu: Menu) {
    return (menu.recettes || []).map(rec => D.recettes[rec.id_recette]);
  }

  static getMenuIngredients(menu: Menu): IngredientOptions[] {
    return (menu.ingredients || []).map(ing => {
      return { ingredient: D.ingredients[ing.id_ingredient], options: ing };
    });
  }

  static getMenuOrRecetteProprietaire(item: Menu | Recette) {
    if (!item.id_proprietaire.Valid) return null;
    return D.utilisateurs[item.id_proprietaire.Int64];
  }

  static getRecetteIngredients(rec: Recette): IngredientOptions[] {
    return (rec.ingredients || []).map(ing => {
      return { ingredient: D.ingredients[ing.id_ingredient], options: ing };
    });
  }
}
