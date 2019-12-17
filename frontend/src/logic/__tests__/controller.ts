import { D } from "../controller";

test("load agenda", async () => {
  await D.loadAgenda();
  expect(D.error).toBeNull();
  expect(D.agenda.sejours).toHaveLength(3);
});

test("ajoute ingredient", async () => {
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
  await D.deleteIngredient(ing, false);
  expect(Object.keys(D.ingredients)).toHaveLength(l);
});
