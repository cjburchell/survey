import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Observable} from "rxjs/index";

export interface Choice{
  id: string
  text: string
}

export interface Question {
  id: string,
  text: string,
  type: string,
  Choices: Choice[]
}

export interface Survey{
  name: string
  id: string
  questions: Question[]
}

export interface Result{
  surveyId: string,
  questionId: string,
  answer: string,
  count: number
}

export interface Answer{
  questionId: string,
  answer: string
}

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
  })
};

@Injectable({
  providedIn: 'root'
})
export class SurveyService {

  constructor(private http: HttpClient) { }

  getSurvey(surveyId :string): Observable<Survey>{
    return this.http.get<Survey>(`/survey/${surveyId}`);
  }

  getResults(surveyId :string): Observable<Result[]>{
    return this.http.get<Result[]>(`/survey/${surveyId}/results`);
  }

  getGetResultsForQuestion(surveyId :string, questionId): Observable<Result[]>{
    return this.http.get<Result[]>(`/survey/${surveyId}/results/${questionId}`);
  }

  setAnswers(surveyId :string, answers :Answer[]){
    return this.http.post<Answer[]>(`/survey/${surveyId}/answers`, answers, httpOptions).pipe()
  }
}
