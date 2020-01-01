import { C } from "../controller";
import { Repas } from "../types";

test("resoud ingredients repas", async () => {
  await C.data.loadAgenda();
  expect(C.notifications.getError()).toBeNull();

  const repas: Repas[] = [];
  C.iterateAllRepas((sej, rep) => {
    repas.push(rep);
  });
  if (repas.length == 0) return;
  const res = await C.calculs.resoudIngredientsRepas(repas[0].id);
  expect(C.notifications.getError()).toBeNull();
});
