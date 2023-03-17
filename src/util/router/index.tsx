import NamespaceLayout from "../../pages/namespace/Layout";
import OnboardingPage from "../../pages/OnboardingPage";
import { createBrowserRouter } from "react-router-dom";
import { pages } from "./pages";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <OnboardingPage />,
  },
  {
    path: "/:namespace",
    element: <NamespaceLayout />,
    children: Object.values(pages).map((page) => page.route),
  },
]);
