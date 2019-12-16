import { Data } from "../controller";

test("load agenda", async () => {
  const d = new Data();
  await d.loadAgenda();
  expect(d.agenda.sejours).toHaveLength(3);
});
