import { Component, OnInit, Input } from '@angular/core';
import {Choice, Question} from "../survey.service";

@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: []
})
export class QuestionComponent implements OnInit {
  @Input() question : Question;
  @Input() number : number;

  ngOnInit() {
  }
}
