import {Component, Input, OnInit} from '@angular/core';
import {Question} from "../survey.service";
import {FormGroup} from "@angular/forms";


function range(start: number, end: number) {
  return  Array.from(Array(end - start + 1), (v, k) => start + k);
}


@Component({
  selector: 'app-question',
  templateUrl: './question.component.html'
})
export class QuestionComponent implements OnInit {
  @Input() group: FormGroup;
  @Input() question : Question;
  @Input() number : number;
  choices: number[];

  ngOnInit() {
    if (this.question.type === "Range"){
      this.choices = range(this.question.choices[0].value, this.question.choices[1].value);
    }
  }
}
