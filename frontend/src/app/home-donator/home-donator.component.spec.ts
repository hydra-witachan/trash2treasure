import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeDonatorComponent } from './home-donator.component';

describe('HomeDonatorComponent', () => {
  let component: HomeDonatorComponent;
  let fixture: ComponentFixture<HomeDonatorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [HomeDonatorComponent]
    });
    fixture = TestBed.createComponent(HomeDonatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
