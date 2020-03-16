import { C } from "../controller";
import { RepasWithGroupe } from "../types";

test("resoud ingredients", async () => {
  await C.data.loadSejours();
  expect(C.notifications.getError()).toBeNull();

  const repas: RepasWithGroupe[] = [];
  C.iterateAllRepas((sej, rep) => {
    repas.push(rep);
  });
  if (repas.length == 0) return;
  await C.calculs.resoudIngredientsRepas(repas[0].id);
  expect(C.notifications.getError()).toBeNull();

  const sejour = Object.values(C.data.sejours).filter(
    s => Object.keys(s.journees).length > 0
  )[0];
  const N = Object.keys(sejour.journees).length;

  const jourOffsets = Object.keys(sejour.journees).map(k => Number(k));
  const res = await C.calculs.resoudIngredientsJournees(
    sejour.sejour.id,
    jourOffsets
  );
  expect(C.notifications.getError()).toBeNull();
  expect(res).not.toBeUndefined();
  if (res == undefined) return;
  expect(res.date_ingredients).toHaveLength(N);
});
