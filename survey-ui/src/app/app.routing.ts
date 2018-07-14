import {RouterModule, Routes} from "@angular/router";
import {SurveyComponent} from "./survey/survey.component";
import {ResultsComponent} from "./results/results.component";


const appRoutes: Routes = [
  {path: 'results/:surveyId', component: ResultsComponent },
  {path: ':surveyId', component: SurveyComponent },
  {path: '', pathMatch:'full', redirectTo: 's1'}
  ];

export const routing = RouterModule.forRoot(appRoutes);
