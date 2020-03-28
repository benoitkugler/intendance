import { C } from "../controller";
import { HoraireFields } from "../enums";
import { toNullableId } from "../types2";

const IdUtilisateur = 2;

test("load agenda", async () => {
  await C.data.loadSejours();
  expect(C.notifications.getError()).toBeNull();
});

test("crud ingredient", async () => {
  await C.data.loadIngredients();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.data.ingredients || {}).length;
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
    unite: "P"
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.ingredients || {})).toHaveLength(l + 1);
  if (!ing) return;

  await C.data.updateIngredient({
    id: ing.id,
    nom:
      "Concombres" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    conditionnement: { unite: "Kg", quantite: 4 },
    callories: {},
    categorie: "nouvelle cat&gori",
    unite: "P"
  });
  expect(C.notifications.getError()).toBeNull();

  await C.data.deleteIngredient(ing.id, false);
  expect(Object.keys(C.data.ingredients || {})).toHaveLength(l);
});

test("crud recette", async () => {
  await Promise.all([C.data.loadRecettes(), C.data.loadIngredients()]);
  expect(C.notifications.getError()).toBeNull();

  const ingId = Number(Object.keys(C.data.ingredients || {})[0]);

  const l = Object.keys(C.data.recettes).length;
  let recette = await C.data.createRecette({
    nom:
      "Gratin de semoule" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    id_utilisateur: { Valid: true, Int64: IdUtilisateur },
    mode_emploi: "BAtter les oeufs en eige....",
    ingredients: [
      {
        id_ingredient: ingId,
        cuisson: "buit",
        quantite: 4
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
      quantite: 4
    }
  ];

  recette = await C.data.updateRecette(recette);
  expect(C.notifications.getError()).toBeNull();
  if (!recette) return;
  expect(recette.ingredients).toHaveLength(1);

  await C.data.deleteRecette(recette.id);
  expect(Object.keys(C.data.recettes)).toHaveLength(l);
});

test("crud menu", async () => {
  await C.data.loadAllMenus();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.data.menus).length;
  let menu = await C.data.createMenu({
    id_utilisateur: { Valid: true, Int64: IdUtilisateur },
    commentaire: "BAtter les oeufs en eige....",
    ingredients: [],
    recettes: []
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.menus)).toHaveLength(l + 1);
  if (!menu) return;

  const ingId = Number(Object.keys(C.data.ingredients || {})[0]);
  const recId = Number(Object.keys(C.data.recettes)[0]);

  menu.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4
    }
  ];
  menu.recettes = [recId];

  menu = await C.data.updateMenu(menu);
  expect(C.notifications.getError()).toBeNull();
  if (!menu) return;
  expect(menu.ingredients).toHaveLength(1);
  expect(menu.recettes).toHaveLength(1);

  await C.data.deleteMenu(menu.id);
  expect(Object.keys(C.data.menus)).toHaveLength(l);
}, 10000);

test("crud sejour", async () => {
  const l = Object.keys(C.data.sejours.sejours || {}).length;
  let sejour = await C.data.createSejour({
    date_debut: new Date().toISOString(),
    nom: "C2 Again !",
    id_utilisateur: IdUtilisateur
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.sejours.sejours || {})).toHaveLength(l + 1);
  if (!sejour) return;

  sejour.nom = "Ah non C3";
  sejour.date_debut = new Date().toISOString();
  sejour = await C.data.updateSejour(sejour);
  expect(C.notifications.getError()).toBeNull();
  if (!sejour) return;
  expect(sejour.nom).toBe("Ah non C3");

  await C.data.deleteSejour(sejour);
  expect(Object.keys(C.data.sejours.sejours || {})).toHaveLength(l);
});

test("crud groupe", async () => {
  await C.data.loadSejours();
  expect(C.notifications.getError()).toBeNull();

  const sejourId = Number(Object.keys(C.data.sejours.sejours || {})[0]);

  const l = Object.keys(C.data.sejours.groupes || {}).length;
  let groupe = await C.data.createGroupe({
    id_sejour: sejourId,
    nom: "Moussaillons",
    couleur: "#787878",
    nb_personnes: 0
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.data.sejours.groupes || {})).toHaveLength(l + 1);
  if (!groupe) return;

  groupe.nom = "Ah non Marins";
  groupe.couleur = "black";
  groupe = await C.data.updateGroupe(groupe);
  expect(C.notifications.getError()).toBeNull();
  if (!groupe) return;
  expect(groupe.nom).toBe("Ah non Marins");

  await C.data.deleteGroupe(groupe);
  expect(Object.keys(C.data.sejours.groupes || {})).toHaveLength(l);
});

test("crud repas", async () => {
  await C.data.loadMenus();
  expect(C.notifications.getError()).toBeNull();

  const menuId = Number(Object.keys(C.data.menus)[0]);
  const sejourId = Number(Object.keys(C.data.sejours.sejours || {})[0]);

  await C.data.createRepas({
    horaire: HoraireFields.Cinquieme,
    id_sejour: sejourId,
    jour_offset: 2,
    offset_personnes: 50,
    groupes: [],
    recettes: [],
    ingredients: []
  });
  expect(C.notifications.getError()).toBeNull();

  const repas = ((C.data.sejours.sejours || {})[sejourId].repass || [])[0];
  repas.horaire = HoraireFields.Midi;
  await C.data.updateManyRepas([repas]);
  expect(C.notifications.getError()).toBeNull();

  await C.data.deleteRepas(repas);
});

test("crud produits", async () => {
  await C.data.loadIngredients();
  expect(C.notifications.getError()).toBeNull();
});
