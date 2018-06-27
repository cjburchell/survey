import {Component, OnInit, Input} from '@angular/core';
import {SurveyService, Survey, Result} from '../survey.service';
import {Router, ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
})
export class ResultsComponent implements OnInit {
  surveyId: string;
  survey: Survey;
  results: Result[];
  submitCount: number;

  constructor(private surveyService: SurveyService, private router: Router, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit() {
    this.activatedRoute.params.subscribe(params=>{
      this.surveyId = params['surveyId'];

      this.surveyService.getSurvey(this.surveyId).subscribe((survey: Survey) => {
        this.survey = survey;
      });

      this.surveyService.getCount(this.surveyId).subscribe(result => {
        this.submitCount = result.count;
      });
    });
  }

  back(){
    this.router.navigate([`/${this.surveyId}`]);
  }
}
