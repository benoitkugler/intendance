import { Controller } from "../controller";
import { metaDev, N } from "../server";
import { Unite, Horaire, Time } from "../api";

const C = new Controller(metaDev, N);

test("load agenda", async () => {
  await C.api.GetSejours();
  expect(C.notifications.getError()).toBeNull();
});

test("crud ingredient", async () => {
  await C.api.GetIngredients();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.api.ingredients || {}).length;
  const ing = await C.api.CreateIngredient({
    nom:
      "Concombres" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    conditionnement: { unite: Unite.Litres, quantite: 2 },
    callories: {},
    categorie: "",
    unite: Unite.Piece
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.api.ingredients || {})).toHaveLength(l + 1);
  if (!ing) return;

  await C.api.UpdateIngredient({
    id: ing.id,
    nom:
      "Concombres" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    conditionnement: { unite: Unite.Kilos, quantite: 4 },
    callories: {},
    categorie: "nouvelle cat&gori",
    unite: Unite.Piece
  });
  expect(C.notifications.getError()).toBeNull();

  await C.api.DeleteIngredient({ id: ing.id, check_produits: false });
  expect(Object.keys(C.api.ingredients || {})).toHaveLength(l);
});

test("crud recette", async () => {
  await Promise.all([C.api.GetRecettes(), C.api.GetIngredients()]);
  expect(C.notifications.getError()).toBeNull();

  const ingId = Number(Object.keys(C.api.ingredients || {})[0]);

  const l = Object.keys(C.api.recettes).length;
  let recette = await C.api.CreateRecette({
    nom:
      "Gratin de semoule" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    id_utilisateur: { Valid: true, Int64: metaDev.idUtilisateur },
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
  expect(Object.keys(C.api.recettes)).toHaveLength(l + 1);
  if (!recette) return;

  recette.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4
    }
  ];

  recette = await C.api.UpdateRecette(recette);
  expect(C.notifications.getError()).toBeNull();
  if (!recette) return;
  expect(recette.ingredients).toHaveLength(1);

  await C.api.DeleteRecette({ id: recette.id });
  expect(Object.keys(C.api.recettes)).toHaveLength(l);
});

test("crud menu", async () => {
  await C.api.loadAllMenus();
  expect(C.notifications.getError()).toBeNull();

  const l = Object.keys(C.api.menus).length;
  let menu = await C.api.CreateMenu({
    id_utilisateur: { Valid: true, Int64: metaDev.idUtilisateur },
    commentaire: "BAtter les oeufs en eige....",
    ingredients: [],
    recettes: []
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.api.menus)).toHaveLength(l + 1);
  if (!menu) return;

  const ingId = Number(Object.keys(C.api.ingredients || {})[0]);

  let recette = await C.api.CreateRecette({
    nom:
      "Gratin de semoule" +
      Math.random()
        .toString(36)
        .replace(/[^a-z]+/g, "")
        .substr(0, 5),
    id_utilisateur: { Valid: true, Int64: metaDev.idUtilisateur },
    ingredients: [],
    mode_emploi: ""
  });
  if (recette === undefined) return;

  menu.ingredients = [
    {
      id_ingredient: ingId,
      cuisson: "buit",
      quantite: 4
    }
  ];
  menu.recettes = [recette.id];

  menu = await C.api.UpdateMenu(menu);
  expect(C.notifications.getError()).toBeNull();
  if (!menu) return;
  expect(menu.ingredients).toHaveLength(1);
  expect(menu.recettes).toHaveLength(1);

  await C.api.DeleteMenu({ id: menu.id });
  expect(Object.keys(C.api.menus)).toHaveLength(l);
}, 10000);

test("crud sejour", async () => {
  const l = Object.keys(C.api.sejours.sejours || {}).length;
  let sejour = await C.api.CreateSejour({
    date_debut: new Date().toISOString() as Time,
    nom: "C2 Again !",
    id_utilisateur: metaDev.idUtilisateur
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.api.sejours.sejours || {})).toHaveLength(l + 1);
  if (!sejour) return;

  sejour.nom = "Ah non C3";
  sejour.date_debut = new Date().toISOString() as Time;
  sejour = await C.api.UpdateSejour(sejour);
  expect(C.notifications.getError()).toBeNull();
  if (!sejour) return;
  expect(sejour.nom).toBe("Ah non C3");

  await C.api.DeleteSejour(sejour);
  expect(Object.keys(C.api.sejours.sejours || {})).toHaveLength(l);
});

test("crud groupe", async () => {
  await C.api.GetSejours();
  expect(C.notifications.getError()).toBeNull();

  const sejourId = Number(Object.keys(C.api.sejours.sejours || {})[0]);

  const l = Object.keys(C.api.sejours.groupes || {}).length;
  let groupe = await C.api.CreateGroupe({
    id_sejour: sejourId,
    nom: "Moussaillons",
    couleur: "#787878",
    nb_personnes: 0
  });
  expect(C.notifications.getError()).toBeNull();
  expect(Object.keys(C.api.sejours.groupes || {})).toHaveLength(l + 1);
  if (!groupe) return;

  groupe.nom = "Ah non Marins";
  groupe.couleur = "black";
  groupe = await C.api.UpdateGroupe(groupe);
  expect(C.notifications.getError()).toBeNull();
  if (!groupe) return;
  expect(groupe.nom).toBe("Ah non Marins");

  await C.api.DeleteGroupe(groupe);
  expect(Object.keys(C.api.sejours.groupes || {})).toHaveLength(l);
});

test("crud repas", async () => {
  await C.api.GetMenus();
  expect(C.notifications.getError()).toBeNull();

  const menuId = Number(Object.keys(C.api.menus)[0]);
  const sejourId = Number(Object.keys(C.api.sejours.sejours || {})[0]);

  await C.api.CreateRepas({
    horaire: Horaire.Cinquieme,
    id_sejour: sejourId,
    jour_offset: 2,
    offset_personnes: 50,
    anticipation: 0,
    groupes: [],
    recettes: [],
    ingredients: []
  });
  expect(C.notifications.getError()).toBeNull();

  const repas = ((C.api.sejours.sejours || {})[sejourId].repass || [])[0];
  repas.horaire = Horaire.Midi;
  await C.api.UpdateManyRepas([repas]);
  expect(C.notifications.getError()).toBeNull();

  await C.api.DeleteRepas(repas);
});

test("crud produits", async () => {
  await C.api.GetIngredients();
  expect(C.notifications.getError()).toBeNull();
});
