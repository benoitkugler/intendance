import { Controller } from "../controller";
import { metaDev, N } from "../server";
import { RepasComplet } from "../api";

const C = new Controller(metaDev, N);

test("resoud ingredients", async () => {
  await C.api.GetSejours();
  expect(C.notifications.getError()).toBeNull();

  const repas: RepasComplet[] = [];
  C.iterateAllRepas((sej, rep) => {
    repas.push(rep);
  });
  if (repas.length == 0) return;
  await C.resoudIngredientsRepas(repas[0].id);
  expect(C.notifications.getError()).toBeNull();

  const sejour = Object.values(C.api.sejours.sejours || {})[0];

  const res = await C.resoudIngredientsJournees(sejour.id, [0, 1, 2, 3]);
  expect(C.notifications.getError()).toBeNull();
  expect(res).not.toBeUndefined();
});
