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

  static getRecetteIngredients(rec: Recette): IngredientOptions[] {
    return (rec.ingredients || []).map(ing => {
      return { ingredient: D.ingredients[ing.id_ingredient], options: ing };
    });
  }
}
