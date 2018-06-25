import {Component, OnInit, Input} from '@angular/core';
import {SurveyService, Question, Result} from '../survey.service';

export interface ChoiceResult {
  answer: string;
  count: number;
}

function range(start: number, end: number) {
  return  Array.from(Array(end - start + 1), (v, k) => start + k);
}

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
})
export class ResultsComponent implements OnInit {
  @Input() question: Question;
  @Input() number: number;
  @Input() submitCount: number;
  @Input() surveyId: string;
  results: ChoiceResult[];

  constructor(private surveyService: SurveyService) {
  }

  ngOnInit() {
    this.results = [];
    if (this.question) {
      this.surveyService.getGetResultsForQuestion(this.surveyId, this.question.id).subscribe((results: Result[]) => {
        const resultMap = new Map<string, number>();
        if (results) {
          results.forEach(result => {
            resultMap[result.answer] = result.count;
          });
        }

        //noinspection TsLint
        switch (this.question.type) {
          case 'Range':
            range(this.question.choices[0].value, this.question.choices[1].value).forEach(value => {
              const number = value.toString();
              let answer = number;
              if (value === this.question.choices[0].value) {
                answer += ' ' + this.question.choices[0].text;
              }

              if (value === this.question.choices[1].value) {
                answer += ' ' + this.question.choices[1].text;
              }

              this.results.push({
                'answer': answer,
                'count': resultMap[number] !== undefined ? resultMap[number] : 0
              });
            });
            break;
          case 'MultipleChoice':
          case 'MultipleSelection':
            this.question.choices.forEach(choice => {
              this.results.push({
                'answer': choice.text,
                'count': resultMap[choice.id] !== undefined ? resultMap[choice.id] : 0
              });
            });
            break;
          case 'Text':
            if (results) {
              results.forEach(result => {
                this.results.push(result);
              });
            }
            break;
        }
      });
    }
  }
}
