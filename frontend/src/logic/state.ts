import { devMode } from "./data";
import { Controller } from "./controller";

export class State {
  private readonly controller: Controller;

  isLoggedIn = devMode; // en dev mode, connection automatique

  idSejour: number | null = null;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  getSejour() {
    if (this.idSejour == null) return null;
    return this.controller.data.getSejour(this.idSejour);
  }
}
