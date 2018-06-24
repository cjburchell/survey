import { Component, Input } from '@angular/core';
import {Question} from "../survey.service";
import {FormGroup} from "@angular/forms";


@Component({
  selector: 'app-question',
  templateUrl: './question.component.html'
})
export class QuestionComponent  {
  @Input() group: FormGroup;
  @Input() question : Question;
  @Input() number : number;
}
