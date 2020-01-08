import { C } from "../controller";
import { Repas } from "../types";

test("resoud ingredients", async () => {
  await C.data.loadAgenda();
  expect(C.notifications.getError()).toBeNull();

  const repas: Repas[] = [];
  C.iterateAllRepas((sej, rep) => {
    repas.push(rep);
  });
  if (repas.length == 0) return;
  await C.calculs.resoudIngredientsRepas(repas[0].id);
  expect(C.notifications.getError()).toBeNull();

  const sejour = Object.values(C.data.agenda.sejours).filter(
    s => Object.keys(s.journees).length > 0
  )[0];
  const N = Object.keys(sejour.journees).length;
  let res = await C.calculs.resoudIngredientsSejour(sejour.sejour.id);
  expect(C.notifications.getError()).toBeNull();
  expect(res).not.toBeUndefined();
  if (res == undefined) return;
  expect(res.date_ingredients).toHaveLength(N);

  const jourOffsets = Object.keys(sejour.journees).map(k => Number(k));
  res = await C.calculs.resoudIngredientsJournees(
    sejour.sejour.id,
    jourOffsets
  );
  expect(C.notifications.getError()).toBeNull();
  expect(res).not.toBeUndefined();
  if (res == undefined) return;
  expect(res.date_ingredients).toHaveLength(N);
});