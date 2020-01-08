import { C } from "../controller";

const IdProprietaire = 2;

test("load agenda", async () => {
  await C.data.loadAgenda();
  expect(C.notifications.getError()).toBeNull();
});

test("crud ingredient", async () => {
  await C.data.loadIngredients();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.data.ingredients).length;
  const ing = await C.data.createIngredient({
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
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.ingredients)).toHaveLength(l + 1);
  if (!ing) return;

  await C.data.updateIngredient({
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
  expect(C.notifications.getError()).toBeNull();

  await C.data.deleteIngredient(ing, false);
  expect(Object.keys(C.data.ingredients)).toHaveLength(l);
});

test("crud recette", async () => {
  await Promise.all([C.data.loadRecettes(), C.data.loadIngredients()]);
  expect(C.notifications.getError()).toBeNull();

  const ingId = Number(Object.keys(C.data.ingredients)[0]);

  const l = Object.keys(C.data.recettes).length;
  let recette = await C.data.createRecette({
    nom:
      "Gratin de semoule" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    id_proprietaire: { Valid: true, Int64: IdProprietaire },
    mode_emploi: "BAtter les oeufs en eige....",
    ingredients: [
      {
        id_ingredient: ingId,
        cuisson: "buit",
        quantite: 4,
        id_recette: -1
      }
    ]
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.recettes)).toHaveLength(l + 1);
  if (!recette) return;

  recette.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4,
      id_recette: recette.id
    }
  ];

  recette = await C.data.updateRecette(recette);
  expect(C.notifications.getError()).toBeNull();
  if (!recette) return;
  expect(recette.ingredients).toHaveLength(1);

  await C.data.deleteRecette(recette);
  expect(Object.keys(C.data.recettes)).toHaveLength(l);
});

test("crud menu", async () => {
  await C.data.loadAllMenus();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.data.menus).length;
  let menu = await C.data.createMenu({
    id_proprietaire: { Valid: true, Int64: IdProprietaire },
    commentaire: "BAtter les oeufs en eige....",
    ingredients: [],
    recettes: []
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.menus)).toHaveLength(l + 1);
  if (!menu) return;

  const ingId = Number(Object.keys(C.data.ingredients)[0]);
  const recId = Number(Object.keys(C.data.recettes)[0]);

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

  menu = await C.data.updateMenu(menu);
  expect(C.notifications.getError()).toBeNull();
  if (!menu) return;
  expect(menu.ingredients).toHaveLength(1);
  expect(menu.recettes).toHaveLength(1);

  await C.data.deleteMenu(menu);
  expect(Object.keys(C.data.menus)).toHaveLength(l);
}, 10000);

test("crud sejour", async () => {
  await C.data.loadMenus();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.data.agenda.sejours).length;
  let sejour = await C.data.createSejour({
    date_debut: new Date().toISOString(),
    nom: "C2 Again !",
    id_proprietaire: IdProprietaire
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.agenda.sejours)).toHaveLength(l + 1);
  if (!sejour) return;

  sejour.nom = "Ah non C3";
  sejour.date_debut = new Date().toISOString();
  sejour = await C.data.updateSejour(sejour);
  expect(C.notifications.getError()).toBeNull();
  if (!sejour) return;
  expect(sejour.nom).toBe("Ah non C3");

  await C.data.deleteSejour(sejour);
  expect(Object.keys(C.data.agenda.sejours)).toHaveLength(l);
});

test("crud repas", async () => {
  await C.data.loadMenus();
  expect(C.notifications.getError()).toBeNull();

  const menuId = Number(Object.keys(C.data.menus)[0]);
  const sejourId = Number(Object.keys(C.data.agenda.sejours)[0]);

  const journee = (C.data.agenda.sejours[sejourId]?.journees || {})[2];
  const l = journee?.menus?.length || 0;
  await C.data.createRepas({
    horaire: { heure: 10, minute: 20 },
    id_menu: menuId,
    id_sejour: sejourId,
    jour_offset: 2,
    nb_personnes: 50
  });
  expect(C.notifications.getError()).toBeNull();
  let menus = C.data.agenda.sejours[sejourId]!.journees[2]?.menus || [];
  expect(menus).toHaveLength(l + 1);

  const repas = menus[0];
  repas.horaire = { heure: 12, minute: 20 };
  await C.data.updateManyRepas([repas]);
  expect(C.notifications.getError()).toBeNull();

  await C.data.deleteRepas(repas);
  menus = C.data.agenda.sejours[sejourId]!.journees[2]?.menus || [];
  expect(menus).toHaveLength(l);
});