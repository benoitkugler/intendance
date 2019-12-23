import { D } from "../controller";
import { NS } from "../notifications";

const IdProprietaire = 2;

test("load agenda", async () => {
  await D.loadAgenda();
  expect(NS.getError()).toBeNull();
});

test("crud ingredient", async () => {
  await D.loadIngredients();
  expect(NS.getError()).toBeNull();

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
  expect(NS.getError()).toBeNull();
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
  expect(NS.getError()).toBeNull();

  await D.deleteIngredient(ing, false);
  expect(Object.keys(D.ingredients)).toHaveLength(l);
});

test("crud recette", async () => {
  await Promise.all([D.loadRecettes(), D.loadIngredients()]);
  expect(NS.getError()).toBeNull();

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
  expect(NS.getError()).toBeNull();
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
  expect(NS.getError()).toBeNull();
  if (!recette) return;
  expect(recette.ingredients).toHaveLength(1);

  await D.deleteRecette(recette);
  expect(Object.keys(D.recettes)).toHaveLength(l);
});

test("crud menu", async () => {
  await Promise.all([D.loadRecettes(), D.loadIngredients(), D.loadMenus()]);
  expect(NS.getError()).toBeNull();

  const l = Object.keys(D.menus).length;
  let menu = await D.createMenu({
    id_proprietaire: { Valid: true, Int64: IdProprietaire },
    commentaire: "BAtter les oeufs en eige....",
    ingredients: [],
    recettes: []
  });
  expect(NS.getError()).toBeNull();
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
  expect(NS.getError()).toBeNull();
  if (!menu) return;
  expect(menu.ingredients).toHaveLength(1);
  expect(menu.recettes).toHaveLength(1);

  await D.deleteMenu(menu);
  expect(Object.keys(D.menus)).toHaveLength(l);
});

test("crud sejour", async () => {
  await D.loadMenus();
  expect(NS.getError()).toBeNull();

  const l = Object.keys(D.agenda.sejours).length;
  let sejour = await D.createSejour({
    date_debut: new Date().toISOString(),
    nom: "C2 Again !",
    id_proprietaire: IdProprietaire
  });
  expect(NS.getError()).toBeNull();
  expect(Object.keys(D.agenda.sejours)).toHaveLength(l + 1);
  if (!sejour) return;

  sejour.nom = "Ah non C3";
  sejour.date_debut = new Date().toISOString();
  sejour = await D.updateSejour(sejour);
  expect(NS.getError()).toBeNull();
  if (!sejour) return;
  expect(sejour.nom).toBe("Ah non C3");

  await D.deleteSejour(sejour);
  expect(Object.keys(D.agenda.sejours)).toHaveLength(l);
});

test("crud repas", async () => {
  await D.loadMenus();
  expect(NS.getError()).toBeNull();

  const menuId = Number(Object.keys(D.menus)[0]);
  const sejourId = Number(Object.keys(D.agenda.sejours)[0]);

  const journee = (D.agenda.sejours[sejourId]?.journees || {})[2];
  const l = journee?.menus?.length || 0;
  await D.createRepas({
    horaire: { heure: 10, minute: 20 },
    id_menu: menuId,
    id_sejour: sejourId,
    jour_offset: 2,
    nb_personnes: 50
  });
  expect(NS.getError()).toBeNull();
  let menus = D.agenda.sejours[sejourId]!.journees[2]?.menus || [];
  expect(menus).toHaveLength(l + 1);

  const repas = menus[0];
  repas.horaire = { heure: 12, minute: 20 };
  await D.updateManyRepas([repas]);
  expect(NS.getError()).toBeNull();

  await D.deleteRepas(repas);
  menus = D.agenda.sejours[sejourId]!.journees[2]?.menus || [];
  expect(menus).toHaveLength(l);
});
