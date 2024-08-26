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
    //window.HSCombobox.autoInit()
  });
  nuxtApp.hook("page:load", () => {
    window.HSStaticMethods.autoInit();
    //window.HSCombobox.autoInit()
  });
  nuxtApp.hook("page:finish", () => {
    window.HSStaticMethods.autoInit();
    //window.HSCombobox.autoInit()
  });
});

//window.HSStaticMethods.autoInit();