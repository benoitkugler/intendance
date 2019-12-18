import { D } from "../controller";

const IdProprietaire = 2;

test("load agenda", async () => {
  await D.loadAgenda();
  expect(D.error).toBeNull();
  expect(D.agenda.sejours).toHaveLength(3);
});

test("crud ingredient", async () => {
  await D.loadIngredients();
  expect(D.error).toBeNull();

  const l = Object.keys(D.ingredients).length;
  const ing = await D.createIngredient({
    nom:
      "Concombres" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    conditionnement: { unite: "L", quantite: 2 },
    callories: {},
    categorie: "",
    unite: "L"
  });
  expect(D.error).toBeNull();
  expect(Object.keys(D.ingredients)).toHaveLength(l + 1);
  if (!ing) return;

  await D.updateIngredient({
    id: ing.id,
    nom:
      "Concombres" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    conditionnement: { unite: "L", quantite: 4 },
    callories: {},
    categorie: "nouvelle cat&gori",
    unite: "Kg"
  });
  expect(D.error).toBeNull();

  await D.deleteIngredient(ing, false);
  expect(Object.keys(D.ingredients)).toHaveLength(l);
});

test("crud recette", async () => {
  await Promise.all([D.loadRecettes(), D.loadIngredients()]);
  expect(D.error).toBeNull();

  const l = Object.keys(D.recettes).length;
  let recette = await D.createRecette({
    nom:
      "Gratin de semoule" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    id_proprietaire: { Valid: true, Int64: IdProprietaire },
    mode_emploi: "BAtter les oeufs en eige....",
    ingredients: []
  });
  expect(D.error).toBeNull();
  expect(Object.keys(D.recettes)).toHaveLength(l + 1);
  if (!recette) return;

  const ingId = Number(Object.keys(D.ingredients)[0]);

  recette.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4,
      id_recette: recette.id
    }
  ];

  recette = await D.updateRecette(recette);
  expect(D.error).toBeNull();
  if (!recette) return;
  expect(recette.ingredients).toHaveLength(1);

  await D.deleteRecette(recette);
  expect(Object.keys(D.recettes)).toHaveLength(l);
});

test("crud menu", async () => {
  await Promise.all([D.loadRecettes(), D.loadIngredients(), D.loadMenus()]);
  expect(D.error).toBeNull();

  const l = Object.keys(D.menus).length;
  let menu = await D.createMenu({
    id_proprietaire: { Valid: true, Int64: IdProprietaire },
    commentaire: "BAtter les oeufs en eige....",
    ingredients: [],
    recettes: []
  });
  expect(D.error).toBeNull();
  expect(Object.keys(D.menus)).toHaveLength(l + 1);
  if (!menu) return;

  const ingId = Number(Object.keys(D.ingredients)[0]);
  const recId = Number(Object.keys(D.recettes)[0]);

  menu.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4,
      id_menu: menu.id
    }
  ];
  menu.recettes = [
    {
      id_recette: recId,
      id_menu: menu.id
    }
  ];

  menu = await D.updateMenu(menu);
  expect(D.error).toBeNull();
  if (!menu) return;
  expect(menu.ingredients).toHaveLength(1);
  expect(menu.recettes).toHaveLength(1);

  await D.deleteMenu(menu);
  expect(Object.keys(D.menus)).toHaveLength(l);
});
