import {Component, OnInit, Input} from '@angular/core';
import {SurveyService, Question, Result, Choice} from "../survey.service";

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
})
export class ResultsComponent implements OnInit {
  @Input() question : Question;
  @Input() number : number;
  @Input() submitCount : number;
  @Input() surveyId : string;
  results: Result[];
  choices: Map<string, Choice>;

  constructor(private surveyService: SurveyService) { }

  ngOnInit() {
    this.choices = new Map<string, Choice>();

    if (this.question){
      if(this.question.choices) {
        this.question.choices.forEach(choice => {
          this.choices[choice.id] = choice;
        });
      }

      this.surveyService.getGetResultsForQuestion(this.surveyId, this.question.id).subscribe((results: Result[])=>{
        this.results = results;
      })
    }
  }
}
