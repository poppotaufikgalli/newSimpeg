import "preline/preline";
import { type IStaticMethods } from "preline/preline";
declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook("app:beforeMount", () => {
    window.HSStaticMethods.autoInit();
  });
  nuxtApp.hook("page:load", () => {
    window.HSStaticMethods.autoInit();
  });
  nuxtApp.hook("page:finish", () => {
    window.HSStaticMethods.autoInit();
  });
});