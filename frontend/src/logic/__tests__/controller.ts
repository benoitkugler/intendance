import { D } from "../controller";

test("load agenda", async () => {
  await D.loadAgenda();
  expect(D.error).toBeNull();
  expect(D.agenda.sejours).toHaveLength(3);
});

test("ajoute ingredient", async () => {
  const l = Object.keys(D.ingredients).length;
  await D.createIngredient({
    nom: "Concombres",
    conditionnement: { unite: "L", quantite: 2 },
    callories: {},
    categorie: "",
    unite: "L",
    id: 0
  });
  expect(D.error).toBeNull();
  expect(Object.keys(D.ingredients)).toHaveLength(l + 1);
});
