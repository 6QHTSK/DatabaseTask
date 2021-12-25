SELECT Course.Cno,Cname,Sno,Grade from Course left join SC on Course.Cno = SC.Cno;

SELECT Sno,Sname,Sage from Student where Sage = ALL(SELECT Sage from Student where Sname='张立');
SELECT Sno,Sname,Sage from Student where Sage in (SELECT Sage from Student where Sname='张立');
SELECT Sno,Sname,Student.Sage from Student,(SELECT Sage from Student where Sname='张立') as t where Student.Sage = t.Sage;

SELECT Student.Sno, Student.Sname from Student,SC where SC.Sno = Student.Sno and SC.Cno = 3 and SC.Grade BETWEEN 80 AND 89;
SELECT Course.Cno, Course.Cname from Course,SC where Course.Cno=SC.Cno and SC.Sno = '200215122';
SELECT Sno, Cno,Grade from SC as SC1 where (Grade + 5) < (SELECT AVG(Grade) from SC as SC2 WHERE SC2.Cno = SC1.Cno);
SELECT Sno, Sname, Sage from Student where Sage < ALL(SELECT Sage from Student where Ssex = '男');
SELECT Sname, Sdept from Student,SC where Student.Sno = SC.Sno and SC.Cno = '2';

UPDATE Student,SC set Student.Sage = Student.Sage + 2 where Student.Sno = SC.Sno and SC.Grade BETWEEN 80 AND 89;
SELECT Student.* from Student,SC WHERE Student.Sno = SC.Sno and SC.Grade BETWEEN 80 AND 89;
/*Recover*/
UPDATE Student,SC set Student.Sage = Student.Sage - 2 where Student.Sno = SC.Sno and SC.Grade BETWEEN 80 AND 89;

SELECT * from Course;
INSERT INTO Course VALUES (8,'C语言',null,1),(9,'人工智能',null,1);
SELECT * from Course;

DELETE FROM Course where Cname = '人工智能';
SELECT * from Course;
/*Recover*/
DELETE FROM Course where Cname = 'C语言';