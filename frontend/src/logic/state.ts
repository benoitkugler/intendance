import { devMode } from "./data";

export class State {
  isLoggedIn = devMode; // en dev mode, connection automatique

  idSejour: number | null = null;
}
