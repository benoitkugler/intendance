import { C } from "../controller";
import { RepasComplet } from "../types";

test("resoud ingredients", async () => {
  await C.api.GetSejours();
  expect(C.notifications.getError()).toBeNull();

  const repas: RepasComplet[] = [];
  C.iterateAllRepas((sej, rep) => {
    repas.push(rep);
  });
  if (repas.length == 0) return;
  await C.calculs.resoudIngredientsRepas(repas[0].id);
  expect(C.notifications.getError()).toBeNull();

  const sejour = Object.values(C.data.sejours.sejours || {})[0];

  const res = await C.calculs.resoudIngredientsJournees(sejour.id, [
    0,
    1,
    2,
    3
  ]);
  expect(C.notifications.getError()).toBeNull();
  expect(res).not.toBeUndefined();
});
