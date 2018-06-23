import {Component, Input, OnInit} from '@angular/core';
import {Result, Survey, SurveyService} from "../survey.service";

@Component({
  selector: 'app-survey',
  templateUrl: './survey.component.html',
  styleUrls: []
})
export class SurveyComponent implements OnInit {
  @Input() surveyId : string;

  constructor(private surveyService: SurveyService) {
  }

  ngOnInit() {
    console.log(this.surveyId);
    this.surveyService.getSurvey(this.surveyId).subscribe((survey: Survey) => this.survey = survey);
  }

  survey : Survey;
  showResults: boolean;
  results: Result[];

  onViewResults(){
    this.surveyService.getResults(this.surveyId).subscribe((results:  Result[]) => this.results = results);
    this.showResults = true;
  }

  onDone(){
    // TODO: Get all the answers
    // TODO: Send them to the server
    // Update and view the results
    this.onViewResults();
  }
}
