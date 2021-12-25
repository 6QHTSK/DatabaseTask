SELECT Sno,Sname,Sage from Student;
SELECT * from Student WHERE Sdept = 'CS';
SELECT Sno,Cno,Grade from SC where Grade >= 90 or Grade < 60;
SELECT Sname,Ssex,Sage from Student where Sage < 19 or Sage > 20;
SELECT Sname,Sdept from Student where Sdept = 'MA' or Sdept = 'IS';
SELECT Cno, Cname, Ccredit from Course where Cname like '%数据%';
SELECT Sno, Cno from SC where Grade is null;
SELECT MAX(Grade) as highScore, MIN(Grade) as lowScore, AVG(Grade) as average from SC where Sno = '200215121';
SELECT Sno, Grade from SC where Cno = '2' order by Grade;
SELECT Sdept, AVG(Sage) as averageAge from Student group by Sdept;